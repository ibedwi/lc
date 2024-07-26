package dev;

import java.util.*;

// Propblem: https://leetcode.com/problems/zigzag-conversion/description
public class ZigzagConversion {
  public String convert(String s, int numRows) {
    if (numRows < 2) {
      return s;
    }
    int currentCounter = 0;
    boolean isIncrementing = true;

    HashMap<Integer, String> map = new HashMap<>();

    for (int i = 0; i < s.length(); i++) {
      // assign initial value if not exist
      if (!map.containsKey(currentCounter)) {
        map.put(currentCounter, "");
      }

      // adding the text to the row
      map.put(currentCounter, map.get(currentCounter) + s.charAt(i));

      // add current counter
      if (currentCounter == numRows - 1) {
        isIncrementing = false;
      }
      if (currentCounter == 0) {
        isIncrementing = true;
      }

      if (isIncrementing) {
        currentCounter++;
      } else {
        currentCounter--;
      }
    }


    StringBuilder totalString = new StringBuilder();

    for (int i = 0; i < numRows; i++) {
      if (map.containsKey(i)) {
        totalString.append(map.get(i));
      }
    }

    return totalString.toString();
  }
}
