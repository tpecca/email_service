# Email Service

This service uses two providers, one of which serves as a backup in case the other one fails.

## Developer's Approach

The developer's approach was to identify the adapters and configure them according to the providers' requirements. The `EmailSender` is set up to use both providers, with logic to use a fallback service if the primary service fails.

## Building and Running the Docker File

To build the Docker file, use the following command:

```bash
docker build . -t email_service_example
```
To run the Docker file, use the following command:
```bash
docker run email_service_example
```