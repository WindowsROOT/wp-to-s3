# Go S3 Uploader

This is a simple Go program to upload all files inside the `uploads/` directory to an AWS S3 bucket, **excluding** the `uploads/` folder itself from the S3 path.

## 📌 Features
- Recursively uploads all files inside `uploads/`
- Retains the folder structure inside S3 but omits `uploads/`
- Uses AWS SDK for Go v2

## 🛠 Installation

1. **Initialize the Go project and install dependencies:**
   ```bash
   go mod init my-s3-upload
   go get github.com/aws/aws-sdk-go-v2
   go get github.com/aws/aws-sdk-go-v2/config
   go get github.com/aws/aws-sdk-go-v2/service/s3
   ```

2. **Create a `.env` file** and add your AWS credentials:
   ```ini
   AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY
   AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
   AWS_REGION=ap-southeast-1
   AWS_S3_BUCKET=your-bucket-name
   ```

## 🚀 Usage

Run the Go script with:
```bash
   go run main.go
```

## 📂 Directory Structure Before Upload
```
uploads/
├── image1.jpg
├── image2.png
├── docs/
│   ├── file1.pdf
│   ├── file2.docx
```

## 📤 Files in S3 After Upload
```
image1.jpg
image2.png
docs/file1.pdf
docs/file2.docx
```

(Note: `uploads/` is **not included** in the S3 path.)

## 📝 License
This project is open-source and available under the MIT License.

