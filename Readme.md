# Setup

Follow these steps to set up your LeetCode session for the project:

1. Log in to LeetCode using your Google account in a web browser.
2. Open the browser's developer tools (usually F12 or right-click and select "Inspect").
3. Go to the "Application" or "Storage" tab.
4. Under "Cookies", find the cookie named "LEETCODE_SESSION".
5. Copy the value of this cookie.

1. Create a new file named `.env` in the root directory of the project.
2. Add the following line to the `.env` file, replacing the placeholder with the value you copied:

   ```
   LEETCODE_SESSION=your_leetcode_session_cookie_value
   ```

**Note:** Make sure to keep your `.env` file private and do not share or commit it to version control systems.

## Project Information

The main Python script for this project was generated in Cursor text editor with the Claude-3.5-Sonnet model.
