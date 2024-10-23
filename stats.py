import os
from tabulate import tabulate
import re

def count_files_in_subfolders(base_path):
    folder_counts = {}
    
    for folder_name in os.listdir(base_path):
        folder_path = os.path.join(base_path, folder_name)
        if os.path.isdir(folder_path):
            file_count = len([f for f in os.listdir(folder_path) if os.path.isfile(os.path.join(folder_path, f))])
            folder_counts[folder_name] = file_count
    
    return folder_counts

def generate_markdown_table(folder_counts):
    # Sort the folders by count in descending order
    sorted_counts = sorted(folder_counts.items(), key=lambda x: x[1], reverse=True)
    
    # Prepare data for tabulate
    table_data = [["Topic", "Count"]] + sorted_counts
    
    # Generate the markdown table
    markdown_table = tabulate(table_data, headers="firstrow", tablefmt="pipe")
    
    return markdown_table

def update_readme(markdown_content):
    readme_path = "README.md"
    if not os.path.exists(readme_path):
        print(f"The file '{readme_path}' does not exist.")
        return

    with open(readme_path, 'r') as file:
        content = file.read()

    # Define the start and end markers for the statistics section
    start_marker = "# LeetCode Solutions Statistics"
    end_marker = "<!-- End of LeetCode Statistics -->"

    # Check if the markers exist, if not, add them
    if start_marker not in content:
        content += f"\n\n{start_marker}\n\n{end_marker}\n"

    # Replace the content between the markers
    pattern = re.compile(f"{re.escape(start_marker)}.*?{re.escape(end_marker)}", re.DOTALL)
    updated_content = pattern.sub(f"{start_marker}\n\n{markdown_content}\n\n{end_marker}", content)

    with open(readme_path, 'w') as file:
        file.write(updated_content)

    print(f"README.md has been updated with the latest statistics.")

def main():
    base_path = "leetcode_solutions"
    
    if not os.path.exists(base_path):
        print(f"The directory '{base_path}' does not exist.")
        return
    
    folder_counts = count_files_in_subfolders(base_path)
    
    if not folder_counts:
        print("No subfolders found in the 'leetcode_solutions' directory.")
        return
    
    markdown_output = generate_markdown_table(folder_counts)
    
    total_files = sum(folder_counts.values())
    markdown_output += f"\n\n**Total number of solutions:** {total_files}\n"
    
    update_readme(markdown_output)

if __name__ == "__main__":
    main()
