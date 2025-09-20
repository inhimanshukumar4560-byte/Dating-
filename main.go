package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Render हमें बताता है कि किस पोर्ट पर सर्वर चलाना है।
	// अगर यह नहीं बताया गया है, तो हम 8080 का उपयोग करेंगे।
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// यह सर्वर को बताता है कि वर्तमान डायरेक्टरी ("./") से फाइलें दिखानी हैं।
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	// सर्वर शुरू करें
	log.Printf("Listening on :%s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
