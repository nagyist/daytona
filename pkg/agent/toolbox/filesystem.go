// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
)

type FileInfo struct {
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	Mode        string `json:"mode"`
	ModTime     string `json:"modTime"`
	IsDir       bool   `json:"isDir"`
	Owner       string `json:"owner"`
	Group       string `json:"group"`
	Permissions string `json:"permissions"`
}

type ReplaceRequest struct {
	Files    []string `json:"files" binding:"required"`
	Pattern  string   `json:"pattern" binding:"required"`
	NewValue string   `json:"newValue" binding:"required"`
}

func getFileInfo(path string) (FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, err
	}

	stat := info.Sys().(*syscall.Stat_t)
	return FileInfo{
		Name:        info.Name(),
		Size:        info.Size(),
		Mode:        info.Mode().String(),
		ModTime:     info.ModTime().String(),
		IsDir:       info.IsDir(),
		Owner:       strconv.FormatUint(uint64(stat.Uid), 10),
		Group:       strconv.FormatUint(uint64(stat.Gid), 10),
		Permissions: fmt.Sprintf("%04o", info.Mode().Perm()),
	}, nil
}

func listFiles(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		path = "."
	}

	files, err := os.ReadDir(path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var fileInfos []FileInfo
	for _, file := range files {
		info, err := getFileInfo(filepath.Join(path, file.Name()))
		if err != nil {
			continue
		}
		fileInfos = append(fileInfos, info)
	}

	c.JSON(200, fileInfos)
}

func uploadFile(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dst := filepath.Join(path, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "file uploaded successfully"})
}

func downloadFile(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}

	c.File(path)
}

func deleteFile(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}

	if err := os.Remove(path); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "file deleted successfully"})
}

func searchFiles(c *gin.Context) {
	path := c.Query("path")
	pattern := c.Query("pattern")
	if path == "" || pattern == "" {
		c.JSON(400, gin.H{"error": "path and pattern are required"})
		return
	}

	var matches []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}
		if matched, _ := filepath.Match(pattern, info.Name()); matched {
			matches = append(matches, path)
		}
		return nil
	})

	if err != nil {
		c.JSON(200, matches)
		return
	}

	c.JSON(200, matches)
}

func getFileDetails(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}

	info, err := getFileInfo(path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, info)
}

func setFilePermissions(c *gin.Context) {
	path := c.Query("path")
	owner := c.Query("owner")
	group := c.Query("group")
	mode := c.Query("mode")

	if path == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}

	if mode != "" {
		modeNum, err := strconv.ParseUint(mode, 8, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid mode format"})
			return
		}
		if err := os.Chmod(path, os.FileMode(modeNum)); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	if owner != "" || group != "" {
		uid := -1
		gid := -1

		if owner != "" {
			uidNum, err := strconv.Atoi(owner)
			if err != nil {
				c.JSON(400, gin.H{"error": "invalid owner format"})
				return
			}
			uid = uidNum
		}

		if group != "" {
			gidNum, err := strconv.Atoi(group)
			if err != nil {
				c.JSON(400, gin.H{"error": "invalid group format"})
				return
			}
			gid = gidNum
		}

		if err := os.Chown(path, uid, gid); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"message": "permissions updated successfully"})
}

func findInFiles(c *gin.Context) {
	path := c.Query("path")
	pattern := c.Query("pattern")
	if path == "" || pattern == "" {
		c.JSON(400, gin.H{"error": "path and pattern are required"})
		return
	}

	type Match struct {
		File    string `json:"file"`
		Line    int    `json:"line"`
		Content string `json:"content"`
	}

	var matches []Match
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		file, err := os.Open(filePath)
		if err != nil {
			return nil
		}
		defer file.Close()

		buf := make([]byte, 512)
		n, _ := file.Read(buf)
		for i := 0; i < n; i++ {
			// Skip binary files
			if buf[i] == 0 {
				return nil
			}
		}

		file.Seek(0, 0)

		scanner := bufio.NewScanner(file)
		lineNum := 1
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), pattern) {
				matches = append(matches, Match{
					File:    filePath,
					Line:    lineNum,
					Content: scanner.Text(),
				})
			}
			lineNum++
		}
		return nil
	})

	if err != nil {
		c.JSON(200, matches)
		return
	}

	c.JSON(200, matches)
}

func replaceInFiles(c *gin.Context) {
	var req ReplaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	type Result struct {
		File    string `json:"file"`
		Success bool   `json:"success"`
		Error   string `json:"error,omitempty"`
	}

	results := make([]Result, 0, len(req.Files))

	for _, filePath := range req.Files {
		content, err := os.ReadFile(filePath)
		if err != nil {
			results = append(results, Result{
				File:    filePath,
				Success: false,
				Error:   err.Error(),
			})
			continue
		}

		newContent := strings.ReplaceAll(string(content), req.Pattern, req.NewValue)

		err = os.WriteFile(filePath, []byte(newContent), 0644)
		if err != nil {
			results = append(results, Result{
				File:    filePath,
				Success: false,
				Error:   err.Error(),
			})
			continue
		}

		results = append(results, Result{
			File:    filePath,
			Success: true,
		})
	}

	c.JSON(200, results)
}
