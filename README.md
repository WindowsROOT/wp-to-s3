# wp-to-s3

📌 How to Use
1️⃣ Install AWS SDK for Go
Run the following commands to set up the Go project and install dependencies:

go mod init my-s3-upload
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3

2️⃣ Create a .env File and add your AWS credentials:
AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY
AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
AWS_REGION=ap-southeast-1
AWS_S3_BUCKET=your-bucket-name

3️⃣ Run the Go Script
go run main.go
