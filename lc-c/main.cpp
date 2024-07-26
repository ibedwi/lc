#include <iostream>

void identify_identical(int values[], int n) {
    int i, j;
    for (i = 0; i < n; ++i) {
        for (j = 0; j < n; ++j) {
            if (values[i] == values[j]) {
                printf("Twin integers found.\n");
                return;
            }
        }
    }
    printf("No two integers are alike.\n");
}

// Signature of identical right
int identical_right(int snow1[], int snow2[], int start) {
    //  Find the length of array
    int snow2_index;

    for (int offset = 0; offset < 6; ++offset) {
        //        snow2_index = start + offset;
        //        if (snow2_index >= 6)
        //            snow2_index = snow2_index - 6;
        //
        //
        //        if (snow1[offset] != snow2[snow2_index])
        //            return 0;
        if (snow1[offset] != snow2[(start + offset) % 6]) {
            return 0;
        }
    }

    return 1;
}

//int array_length(int arr[]) {
//    int arrSize = sizeof(arr) / sizeof(arr[0]);
//    return arrSize;
//}

int main() {
    int a[5] = {1, 2, 3, 4, 5};
//    identify_identical(a, 5);

    std::cout << "array length: " << sizeof(a) / sizeof(a[0]) << std::endl;
    return 0;
}
