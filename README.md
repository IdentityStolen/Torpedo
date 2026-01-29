# 🚀 Torpedo

Fire torpedoes concurrently & stop as soon as the first enemy ship is destroyed!

## 🎯 Overview

Handcrafted Go program demonstrates concurrent torpedo firing using goroutines and channels. Multiple torpedoes are launched simultaneously, and the program terminates immediately upon the first successful hit.

## 🏃 Running the Program

```bash
go run -race .
```

The `-race` flag enables Go's race detector to ensure thread-safe concurrent operations.

## 📊 Example Outputs

### ✅ Success Scenario

When a target is hit:

```
Miss at {5 8}
Miss at {8 8}
Miss at {2 8}
Miss at {10 8}
Miss at {4 8}
Miss at {6 8}
Miss at {7 8}
Hit at {1 8}
```

### ❌ Failure Scenario

When no targets are hit:

```
Miss at {5 8}
Miss at {4 8}
Miss at {7 8}
Miss at {6 8}
Miss at {2 8}
Miss at {8 8}
Miss at {10 8}
Miss at {9 8}
Miss at {3 8}
No enemy ship destroyed, pray for your life now!
```

## 🔧 Features

- **Concurrent Execution**: Uses Go goroutines for parallel torpedo firing
- **Race-Free**: Designed to pass Go's race detector checks
- **Immediate Termination**: Stops as soon as the first hit is detected
- **Random Delays**: Simulates realistic torpedo travel times