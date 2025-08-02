FROM golang

# Set working directory
WORKDIR /app

# Expose port if needed (optional for CLI tool)
EXPOSE 8000

# Start in interactive shell
CMD ["bash"]
