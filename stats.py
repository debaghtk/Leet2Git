import os
from tabulate import tabulate

def count_files_in_subfolders(base_path):
    folder_counts = {}
    
    for folder_name in os.listdir(base_path):
        folder_path = os.path.join(base_path, folder_name)
        if os.path.isdir(folder_path):
            file_count = len([f for f in os.listdir(folder_path) if os.path.isfile(os.path.join(folder_path, f))])
            folder_counts[folder_name] = file_count
    
    return folder_counts

def print_table(folder_counts):
    # Sort the folders by count in descending order
    sorted_counts = sorted(folder_counts.items(), key=lambda x: x[1], reverse=True)
    
    # Prepare data for tabulate
    table_data = [["Topic", "Count"]] + sorted_counts
    
    # Print the table
    print(tabulate(table_data, headers="firstrow", tablefmt="grid"))

def main():
    base_path = "leetcode_solutions"
    
    if not os.path.exists(base_path):
        print(f"The directory '{base_path}' does not exist.")
        return
    
    folder_counts = count_files_in_subfolders(base_path)
    
    if not folder_counts:
        print("No subfolders found in the 'leetcode_solutions' directory.")
        return
    
    print_table(folder_counts)
    
    total_files = sum(folder_counts.values())
    print(f"\nTotal number of solutions: {total_files}")

if __name__ == "__main__":
    main()

