---

# InfraLive Demo

## Overview
This project provides a simple web interface to deploy an S3 bucket using AWS credentials. It consists of a **backend** written in Go and a **frontend** built with HTML and JavaScript.

## How to Use

### Prerequisites
1. **Go**: Ensure Go is installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2. **AWS Credentials**: You need valid AWS Access Key ID and Secret Access Key with permissions to create S3 buckets.

### Steps to Run the Application

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd InfraLive-demo
   ```

2. **Run the Backend**:
   - Start the Go backend server:
     ```bash
     go run main.go
     ```
   - The server will run on `http://localhost:8080`.

3. **Open the Frontend**:
   - Open the `index.html` file in your browser (e.g., double-click the file or use a local server like `python3 -m http.server`).

4. **Deploy an S3 Bucket**:
   - Enter your AWS credentials, bucket name, and region in the form.
   - Click the **Deploy** button to create the S3 bucket.

   ![Demo Screenshot](https://github.com/user-attachments/assets/0b1ba506-a079-4e1f-a34f-7b2aa6e2e282)

5. **Verify the Bucket**:
   - Use the AWS Management Console or AWS CLI to verify that the bucket has been created:
     ```bash
     aws s3api head-bucket --bucket <bucket-name>
     ```

## Future Enhancements
This is just the beginning! Future updates may include support for deploying additional AWS resources through the same UI.

## Contributing
Contributions are welcome! If you'd like to contribute, please fork the repository and submit a pull request.


