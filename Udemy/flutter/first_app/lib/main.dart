import 'package:flutter/material.dart';
import 'gradient_container.dart';

void main(List<String> args) {
  runApp(
    const MaterialApp(
      home: Scaffold(
        body: GradientContainer(color1: Colors.green, color2: Colors.red),
      ),
    ),
  );
}
