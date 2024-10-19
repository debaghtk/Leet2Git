import os
import requests
import json
from time import sleep
from dotenv import load_dotenv
import logging

# Load environment variables
load_dotenv()

# LeetCode session cookie from environment variable
LEETCODE_SESSION = os.getenv('LEETCODE_SESSION')

# Create a session to maintain cookies
session = requests.Session()
session.cookies.set("LEETCODE_SESSION", LEETCODE_SESSION, domain=".leetcode.com")

# Set up logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

def get_solved_problems():
    url = "https://leetcode.com/graphql"
    query = """
    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
      problemsetQuestionList: questionList(
        categorySlug: $categorySlug
        limit: $limit
        skip: $skip
        filters: $filters
      ) {
        total: totalNum
        questions: data {
          acRate
          difficulty
          freqBar
          frontendQuestionId: questionFrontendId
          isFavor
          paidOnly: isPaidOnly
          status
          title
          titleSlug
          topicTags {
            name
            id
            slug
          }
          hasSolution
          hasVideoSolution
        }
      }
    }
    """
    variables = {
        "categorySlug": "",
        "skip": 0,
        "limit": 5000,  # Adjust if you have solved more than 5000 problems
        "filters": {"status": "AC"}
    }
    
    response = session.post(url, json={"query": query, "variables": variables})
    data = response.json()
    
    if 'errors' in data:
        print(f"Error fetching solved problems: {data['errors']}")
        return []
    
    if 'data' not in data:
        print(f"Unexpected response format. Response: {data}")
        return []
    
    return data['data']['problemsetQuestionList']['questions']

def get_solution(title_slug):
    # First, get the submission ID
    url = "https://leetcode.com/graphql"
    query = """
    query submissionList($offset: Int!, $limit: Int!, $lastKey: String, $questionSlug: String!) {
      submissionList(offset: $offset, limit: $limit, lastKey: $lastKey, questionSlug: $questionSlug) {
        lastKey
        hasNext
        submissions {
          id
          statusDisplay
          lang
        }
      }
    }
    """
    variables = {
        "questionSlug": title_slug,
        "offset": 0,
        "limit": 1,
        "lastKey": None
    }
    
    response = session.post(url, json={"query": query, "variables": variables})
    data = response.json()
    
    if 'errors' in data:
        logging.error(f"Error fetching submission ID for {title_slug}: {data['errors']}")
        return None
    
    if 'data' not in data or 'submissionList' not in data['data'] or 'submissions' not in data['data']['submissionList']:
        logging.error(f"Unexpected response format for {title_slug}. Response: {data}")
        return None
    
    submissions = data['data']['submissionList']['submissions']
    if not submissions:
        logging.error(f"No submissions found for {title_slug}")
        return None
    
    submission_id = submissions[0]['id']
    
    # Now, get the actual code using the submission ID
    query = """
    query submissionDetails($submissionId: Int!) {
      submissionDetails(submissionId: $submissionId) {
        code
        lang {
          name
          verboseName
        }
      }
    }
    """
    variables = {
        "submissionId": int(submission_id)
    }
    
    response = session.post(url, json={"query": query, "variables": variables})
    data = response.json()
    
    if 'errors' in data:
        logging.error(f"Error fetching code for submission {submission_id}: {data['errors']}")
        return None, None
    
    if 'data' not in data or 'submissionDetails' not in data['data']:
        logging.error(f"Unexpected response format for submission {submission_id}. Response: {data}")
        return None, None
    
    submission_details = data['data']['submissionDetails']
    code = submission_details['code']
    lang = submission_details['lang']['name']
    
    logging.info(f"Successfully fetched {lang} solution for {title_slug}")
    return code, lang

def get_file_extension(lang):
    extension_map = {
        'python': 'py',
        'python3': 'py',
        'java': 'java',
        'c++': 'cpp',
        'c': 'c',
        'javascript': 'js',
        'typescript': 'ts',
        'ruby': 'rb',
        'swift': 'swift',
        'go': 'go',
        'golang': 'go',  # Added golang explicitly
        'scala': 'scala',
        'kotlin': 'kt',
        'rust': 'rs',
        'php': 'php',
    }
    lang_lower = lang.lower()
    if lang_lower in extension_map:
        return extension_map[lang_lower]
    elif lang_lower.startswith('c++'):
        return 'cpp'
    elif lang_lower.startswith('c#'):
        return 'cs'
    else:
        logging.warning(f"Unknown language: {lang}. Using .txt extension.")
        return 'txt'

def save_solution(problem, solution, lang):
    folder_name = "leetcode_solutions"
    if not os.path.exists(folder_name):
        os.makedirs(folder_name)
    
    file_extension = get_file_extension(lang)
    file_name = f"{problem['frontendQuestionId']}_{problem['titleSlug']}.{file_extension}"
    file_path = os.path.join(folder_name, file_name)
    
    logging.info(f"Saving solution for {problem['title']} with extension .{file_extension}")
    
    with open(file_path, 'w', encoding='utf-8') as f:
        f.write(f"# {problem['title']}\n")
        f.write(f"# Difficulty: {problem['difficulty']}\n")
        f.write(f"# Language: {lang}\n")
        f.write(f"# Link: https://leetcode.com/problems/{problem['titleSlug']}/\n\n")
        f.write(solution)

def solution_exists(problem, lang):
    folder_name = "leetcode_solutions"
    file_extension = get_file_extension(lang)
    file_name = f"{problem['frontendQuestionId']}_{problem['titleSlug']}.{file_extension}"
    file_path = os.path.join(folder_name, file_name)
    return os.path.exists(file_path)

def main():
    if not LEETCODE_SESSION:
        logging.error("Please set LEETCODE_SESSION in your .env file")
        return

    # Check if we're logged in
    response = session.get("https://leetcode.com/api/problems/all/")
    if response.status_code != 200:
        logging.error(f"Failed to authenticate. Status code: {response.status_code}")
        logging.error("Please check your LEETCODE_SESSION cookie and make sure it's still valid.")
        return

    solved_problems = get_solved_problems()
    
    if not solved_problems:
        logging.warning("No solved problems found. Please check your LEETCODE_SESSION cookie.")
        return

    for problem in solved_problems:
        logging.info(f"Processing: {problem['title']}")
        
        # First, get the language of the solution without fetching the full solution
        lang = get_solution_language(problem['titleSlug'])
        
        if not lang:
            logging.warning(f"Could not determine language for: {problem['title']}")
            continue
        
        if solution_exists(problem, lang):
            logging.info(f"Solution already exists for: {problem['title']} ({lang}). Skipping.")
            continue
        
        logging.info(f"Fetching solution for: {problem['title']}")
        solution, _ = get_solution(problem['titleSlug'])  # We already know the language
        if solution:
            save_solution(problem, solution, lang)
            logging.info(f"Saved {lang} solution for: {problem['title']}")
        else:
            logging.warning(f"Could not fetch solution for: {problem['title']}")
        sleep(1)  # To avoid overwhelming the server

    logging.info(f"\nAll new solutions have been saved to the 'leetcode_solutions' folder.")

# New function to get just the language of the solution
def get_solution_language(title_slug):
    url = "https://leetcode.com/graphql"
    query = """
    query submissionList($offset: Int!, $limit: Int!, $lastKey: String, $questionSlug: String!) {
      submissionList(offset: $offset, limit: $limit, lastKey: $lastKey, questionSlug: $questionSlug) {
        submissions {
          lang
        }
      }
    }
    """
    variables = {
        "questionSlug": title_slug,
        "offset": 0,
        "limit": 1,
        "lastKey": None
    }
    
    response = session.post(url, json={"query": query, "variables": variables})
    data = response.json()
    
    if 'errors' in data:
        logging.error(f"Error fetching language for {title_slug}: {data['errors']}")
        return None
    
    if 'data' not in data or 'submissionList' not in data['data'] or 'submissions' not in data['data']['submissionList']:
        logging.error(f"Unexpected response format for {title_slug}. Response: {data}")
        return None
    
    submissions = data['data']['submissionList']['submissions']
    if not submissions:
        logging.warning(f"No submissions found for {title_slug}")
        return None
    
    return submissions[0]['lang']

if __name__ == "__main__":
    main()
