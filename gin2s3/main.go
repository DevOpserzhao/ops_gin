package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	filename := header.Filename

	//filepath := "http://localhost:8080/file/" + filename
	//c.JSON(http.StatusOK, gin.H{"filepath": filepath})
	bucket := "testzxf"
	//sess, err := session.NewSession(&aws.Config{
	//	Region: aws.String("cn-northwest-1")},
	//)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewSharedCredentials(".aws/credentials", "default"),
	})

	// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
	// for more information on configuring part size, and concurrency.
	//
	// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
	uploader := s3manager.NewUploader(sess)

	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename.
	s3_re := &s3manager.UploadOutput{}

	s3_re, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(filename),
		//Key: aws.String(key),
		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body:        file,
		ContentType: aws.String("image/jpeg"),
		//Metadata:   map[string]*string{
		//	"Content-Type": aws.String("image/jpeg"),
		//},

	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
		c.JSON(http.StatusOK, gin.H{"错误": 000000})
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"s3": s3_re,
		"sever": filepath,
	})

}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "select_file.html", gin.H{})
	})

	router.POST("/upload", upload)

	router.StaticFS("/file", http.Dir("public"))
	router.Run(":8080")

}
