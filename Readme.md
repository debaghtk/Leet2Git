# Setup

Follow these steps to set up your LeetCode session for the project:

```
pip install -r requirements.txt
```

* Log in to LeetCode using your Google account in a web browser.
* Open the browser's developer tools (usually F12 or right-click and select "Inspect").
* Go to the "Application" or "Storage" tab.
* Under "Cookies", find the cookie named "LEETCODE_SESSION".
* Copy the value of this cookie.

* Create a new file named `.env` in the root directory of the project.
* Add the following line to the `.env` file, replacing the placeholder with the value you copied:

   ```
   LEETCODE_SESSION=your_leetcode_session_cookie_value
   ```

**Note:** Make sure to keep your `.env` file private and do not share or commit it to version control systems.

## Project Information

The main Python script for this project was generated in Cursor text editor with the Claude-3.5-Sonnet model.


# LeetCode Solutions Statistics

| Topic                 | Count |
| :-------------------- | ----: |
| Tree                  |    22 |
| Array                 |    22 |
| Graph                 |     7 |
| String                |     6 |
| Hash Table            |     6 |
| Math                  |     5 |
| Trie                  |     4 |
| Linked List           |     4 |
| Matrix                |     3 |
| Divide and Conquer    |     2 |
| Stack                 |     2 |
| Heap (Priority Queue) |     2 |

**Total number of solutions:** 85


<!-- End of LeetCode Statistics -->
