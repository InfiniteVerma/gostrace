#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <sstream>
#include <unistd.h> // For sleep

using namespace std;

// Function to read lines from the file
vector<string> read_lines(const string& filename) {
    ifstream file(filename);
    vector<string> lines;
    string line;

    while (getline(file, line)) {
        lines.push_back(line);
    }
    
    file.close(); // Close the file after reading
    return lines;
}

// Function to write lines back to the file
void write_lines(const string& filename, const vector<string>& lines) {
    ofstream file(filename, ios::trunc); // Truncate the file
    for (const auto& line : lines) {
        file << line << endl;
    }
    
    file.close(); // Close the file after writing
}

int main() {
    const string filename = "data.txt";

    // Infinite loop to read and write to the file
    while (true) {
        // Read lines from the file
        vector<string> lines = read_lines(filename);

        // If there are lines to read
        if (!lines.empty()) {
            // Print the first line and delete it
            cout << "Read: " << lines.front() << endl;
            lines.erase(lines.begin()); // Remove the first line

            // Write remaining lines back to the file
            write_lines(filename, lines);
        } else {
            cout << "No lines to read. Waiting..." << endl;
        }

        // Write a new line to the file
        ofstream file(filename, ios::app); // Append mode
        file << "New line added at " << time(nullptr) << endl;
        file.close(); // Close the file after appending

        // Sleep for a while before the next iteration
        sleep(1);
    }

    return 0;
}

