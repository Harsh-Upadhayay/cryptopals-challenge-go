FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Removed in favour of mount in docker-compose.
# COPY . .

RUN go mod init cryptopals-challenge-go


# Default command (opens a shell for interactive development)
CMD ["bash"]
