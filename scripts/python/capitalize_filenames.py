import os

def rename_files_in_folders(directory):
    for file_name in os.listdir(directory):
        file_path = os.path.join(directory, file_name)
        if os.path.isfile(file_path):
            new_file_name = file_name.capitalize()
            new_file_path = os.path.join(directory, new_file_name)

            if file_path != new_file_path:  # Check if the name needs to change
                os.rename(file_path, new_file_path)
                print(f'Renamed: {file_name} -> {new_file_name} in {file_path}')

rename_files_in_folders('C:\Projects\Go\pokemon\data\imgs')
