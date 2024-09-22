#include <unistd.h>
#include <sys/syscall.h>
#include <cstring>

// Function to write to a file descriptor
ssize_t my_write(int fd, const void *buf, size_t count) {
    return syscall(SYS_write, fd, buf, count);
}

// Function to read from a file descriptor
ssize_t my_read(int fd, void *buf, size_t count) {
    return syscall(SYS_read, fd, buf, count);
}

// Function to exit a process
void my_exit(int status) {
    syscall(SYS_exit, status);
}

int main() {
    // Example usage
    const char *message = "Hello from syscalls!\n";
    my_write(STDOUT_FILENO, message, strlen(message));

    char buffer[100];
    my_read(STDIN_FILENO, buffer, sizeof(buffer)); // Read from standard input
    my_write(STDOUT_FILENO, buffer, strlen(buffer)); // Echo back

    my_exit(0); // Exit the program
}

