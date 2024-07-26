package dev;

import org.junit.Test;

import static org.junit.Assert.*;

public class ZigzagConversionTest {
  @Test
  public void test() {
    // Arrange
    ZigzagConversion zConv = new ZigzagConversion();
    String str = "PAYPALISHIRING";

    // Act
    String result1 = zConv.convert(str, 3 );
    String result2 = zConv.convert(str, 4);
    String result3 = zConv.convert("A", 2);
    String result4 = zConv.convert("ABCD", 2);
    String result5 = zConv.convert("AB", 1);

    // Assert
    assertEquals("PAHNAPLSIIGYIR", result1);
    assertEquals("PINALSIGYAHRPI", result2);
    assertEquals("A", result3);
    assertEquals("ACBD", result4);
    assertEquals("AB", result5);


  }
}