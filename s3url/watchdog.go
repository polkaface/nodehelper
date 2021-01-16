package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cz-theng/czkit-go/log"
)

var files map[string]string

func check_s3_urls() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	bucket := "polkaface-node-db-mirror"
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		log.Error("Unable to list items in bucket  %v", err)
	}

	now := time.Now()
	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		if now.Sub(*item.LastModified).Hours() > 72 {
			log.Info("delete file:%s", *item.Key)
			_, err = svc.DeleteObject(&s3.DeleteObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(*item.Key)})
			if err != nil {
				log.Error("delete file:%s error:%s", *item.Key, err.Error())
			}
			if _, ok := files[*item.Key]; ok {
				delete(files, *item.Key)
			}
			continue
		}

		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(*item.Key),
		})
		urlStr, err := req.Presign(72 * time.Hour)
		if err != nil {
			log.Error("Failed to sign request:", err)
			continue
		}

		log.Info("[%s]:%s", *item.Key, urlStr)
		files[*item.Key] = urlStr
	}
}

func do_timer() {
	log.Info("do_timer")
	now := time.Now()
	log.Info("hour:%d", now.Hour())
	if now.Hour() == 1 {
		check_s3_urls()
	}
}

func watchdog() {
	files = make(map[string]string, 1024)
	check_s3_urls()
	ticker := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-ticker.C:
			do_timer()
		}
	}
}
