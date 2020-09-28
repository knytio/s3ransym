# s3ransym

AWS S3 Ransomware Simulator.

Mimic attackers with access to AWS resources to encrypt S3 buckets for Ransomware

The tool simulates attacker ransomware attacker by

1. Create test S3 bucket
2. Copy files from local folder to S3 bucket
3. Encrypt S3 bucket, it encrypts the files and deletes files in folder

Note: Running the tool on production S3 bucket will encrupt the bucket. Use the key in s3_encrypt.go to decrypt the files.

# Validate Security Polices

1. If you are running from endpoint, check if any endpoint security tools detect or prevent encruption of s3 buckets
2. Check if AWS monitoring tools detect encryption of s3 bucket

# Running

s3ransym

AWS S3 Ransomware Simulator

Usage:
  s3ransym [command]

Available Commands:
  help        Help about any command
  s3          S3 Ransomware Simulator

Flags:
  -h, --help   help for s3ransym

Use "s3ransym [command] --help" for more information about a command.


# Create s3 bucket
s3ransym s3 mb -b bucketname


# Copy files to s3 bucket from local folder "test"
s3ransym s3 cp -b -s bucketname test/

# Encrypt s3 bucket
s3ransym s3 encrypt -b bucketname

