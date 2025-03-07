/**
 * Copyright 2025 Su Yang (soulteary)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/draw"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program <input.png>")
		return
	}

	inputPath := os.Args[1]
	if !strings.HasSuffix(strings.ToLower(inputPath), ".png") {
		fmt.Println("Error: Input file must be a PNG image")
		return
	}

	// 读取源PNG图片
	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inputFile.Close()

	// 解码PNG图片
	srcImg, err := png.Decode(inputFile)
	if err != nil {
		fmt.Printf("Error decoding PNG: %v\n", err)
		return
	}

	// 计算等比缩放后的尺寸
	srcBounds := srcImg.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()

	targetWidth := 640
	targetHeight := 320

	// 计算缩放比例
	widthRatio := float64(targetWidth) / float64(srcWidth)
	heightRatio := float64(targetHeight) / float64(srcHeight)
	ratio := widthRatio

	if heightRatio < widthRatio {
		ratio = heightRatio
	}

	newWidth := int(float64(srcWidth) * ratio)
	newHeight := int(float64(srcHeight) * ratio)

	// 创建目标图片
	dstImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// 使用高质量的缩放算法
	draw.CatmullRom.Scale(dstImg, dstImg.Bounds(), srcImg, srcBounds, draw.Over, nil)

	// 生成输出文件名
	outputPath := filepath.Join(filepath.Dir(inputPath), "bootup.bmp")
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// 编码为BMP格式并保存
	err = bmp.Encode(outputFile, dstImg)
	if err != nil {
		fmt.Printf("Error encoding BMP: %v\n", err)
		return
	}

	fmt.Printf("Successfully converted and resized image to: %s\n", outputPath)
	fmt.Printf("New dimensions: %dx%d\n", newWidth, newHeight)
}
