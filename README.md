
# Selenium Email Sender 📧🚀

Welcome to the **Selenium Email Sender** project! This tool allows you to send emails using **Outlook** (Gmail functionality coming soon!) by automating browser interactions with **Selenium**. This makes it easy to send emails programmatically without using an SMTP server. 

## Features ✨

- Send emails using **Outlook** with ease
- Handles email authentication, subject, and message content
- Supports multiple recipients from a text file
- A simple command-line interface for quick email sending

---

## Prerequisites 🛠

To get started, you'll need the following:

- **Go** (Go 1.17+)
- **Selenium WebDriver** (with **chromedriver** for Chrome)
- **Go Selenium package**: Install it using:
  ```bash
  go get github.com/tebeka/selenium
  ```
- **Chromedriver**: Make sure the `chromedriver` binary is available on your system and the path is set properly.

---

## How to Use 📜

### Step 1: Build the Project 🏗️

Clone this repository to your local machine and navigate into the project directory:

```bash
git clone https://github.com/svg-rs/SeleniumEmailSender.git
cd SeleniumEmailSender
```

Build the project:

```bash
go build -o emailSender main.go
```

---

### Step 2: Run the Program 🚀

Use the command-line interface to send emails. Here's the basic syntax:

```bash
./emailSender -outlook -message "Your message here" -subject "Your subject here" -username "your-email@example.com" -password "your-password" -recipients "recipients.txt"
```

Where:
- `-outlook`: Flag to use Outlook for sending emails.
- `-gmail`: Flag (not yet implemented) for Gmail support.
- `-message`: The content of the email.
- `-subject`: The subject line of the email.
- `-username`: The email address you're sending from.
- `-password`: The password for the email account.
- `-recipients`: A file containing comma-separated email addresses of the recipients.

Example:

```bash
./emailSender -outlook -message "Hello, World!" -subject "Test Email" -username "myemail@example.com" -password "mypassword" -recipients "recipients.txt"
```

---

## Recipients File 📄

The `-recipients` flag requires a file with a list of comma-separated or semicolon-separated email addresses. 

Example `recipients.txt` file format:
```
example1@email.com, example2@email.com, example3@email.com
```

---

## Error Handling 🛑

If any error occurs during the process (e.g., wrong credentials or invalid recipient list), the program will log an error message and stop execution.

---

### Example Output:

```text
[INFO] 2025/03/02 12:30:00 Starting Outlook chromedriver...
[INFO] 2025/03/02 12:30:10 Outlook chromedriver started!
[INFO] 2025/03/02 12:30:15 Outlook opened!
[INFO] 2025/03/02 12:30:20 Logging in...
...
[INFO] 2025/03/02 12:32:50 Email sent successfully to 3 recipients!
```

---

## Contributing 🤝

Feel free to open issues or submit pull requests to enhance the functionality! Any contributions are welcome!

---

## License 📜

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Help Command ❓

If you need help with using the tool, run the following command:

```bash
./emailSender -help
```

It will display all available flags and options.
