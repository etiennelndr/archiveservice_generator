/**
 * MIT License
 *
 * Copyright (c) 2018 CNES
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package utils

import (
	"os"

	"github.com/etiennelndr/archiveservice_generator/constants"
)

// WriteLicense is used to write the License in a specific file
func WriteLicense(file *os.File) error {
	for i := 0; i < len(constants.License); i++ {
		_, err := file.Write([]byte(constants.License[i]))
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteHeader writes the header of a file (License + package name)
func WriteHeader(file *os.File, packageName string) error {
	err := WriteLicense(file)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte("package " + packageName + "\n"))
	if err != nil {
		return err
	}

	return nil
}
