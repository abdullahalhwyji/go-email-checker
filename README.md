# go-email-checker

`go-email-checker` is a simple command-line tool written in Go that checks for the presence of MX, SPF, and DMARC records for a given domain. It helps determine the email configuration and security measures of the domain.

## Features

- Checks for MX (Mail Exchange) records
- Checks for SPF (Sender Policy Framework) records
- Checks for DMARC (Domain-based Message Authentication, Reporting & Conformance) records
- Displays the SPF and DMARC records if available

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/abdullahalhwyji/go-email-checker.git
   cd go-email-checker
   ```

2. **Install Go** (if not already installed)

   Follow the official instructions to install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

3. **Build the application**

   ```bash
   go build -o go-email-checker main.go
   ```

## Usage

Run the application from the command line:

```bash
./go-email-checker
```

Then, enter the domain names you want to check, one per line:

```
example.com
google.com
yahoo.com
```

Press `Ctrl+D` (Linux/Mac) or `Ctrl+Z` (Windows) to end the input.

## Output

The output is a formatted table that shows the domain, presence of MX, SPF, and DMARC records, and the actual SPF and DMARC records if they exist:

```
Domain               hasMX      hasSPF     SPF Record                                          hasDMARC   DMARC Record
--------------------------------------------------------------------------------------------------------------
example.com          true       true       v=spf1 include:_spf.google.com ~all                 true       v=DMARC1; p=none; rua=mailto:dmarc-reports@example.com
```

## Error Handling

If there is an error looking up a record, the tool will log the error message and continue processing the next domain.

