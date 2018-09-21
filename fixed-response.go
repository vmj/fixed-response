/*
 * fixed-response
 * Copyright (C) 2018  <mikko@varri.fi>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4040", "Listening address")
	statusCode := flag.Int("status", 404, "HTTP status code")
	contentType := flag.String("type", "", "Content type")
	filename := flag.String("content", "", "File to send")

	flag.Parse()

	var content []byte
	var contentLength string
	var err error

	if *filename != "" {
		content, err = ioutil.ReadFile(*filename)
		if err != nil {
			log.Fatal(err)
		}
		contentLength = fmt.Sprintf("%d", len(content))
	}
	hasContent := len(content) > 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if *contentType != "" {
			w.Header().Set("Content-Type", *contentType)
		}
		if hasContent {
			w.Header().Set("Content-Length", contentLength)
		}
		w.WriteHeader(*statusCode)

		if hasContent {
			_, err = w.Write(content)
			if err != nil {
				log.Println("Write error:", err)
			}
		}

		log.Printf("%s %s %s (%s)\n", r.RemoteAddr, r.Method, r.URL.Path, r.UserAgent())
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
