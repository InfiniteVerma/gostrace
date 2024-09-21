#include <bits/stdc++.h>
#include <chrono>
#include <thread>
using namespace std;
int main() {
    cout << "Hello\n";
    bool t = false;
    while(1) {
        if(t)
            cout << "Hello\n";
        else
            cout << "World\n";
        fflush(stdout);
        t = !t;
        this_thread::sleep_for(chrono::milliseconds(500));
    }
}
