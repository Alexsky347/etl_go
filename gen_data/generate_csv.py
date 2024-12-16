import csv

# Define the number of lines
num_lines = 1000000

# Open the file in write mode
with open('data.csv', mode='w', newline='') as file:
    writer = csv.writer(file, delimiter=';')
    
    # Write the header
    writer.writerow(['client_id', 'num1', 'num2', 'num3', 'filiale_id'])
    
    # Write the data rows
    for i in range(1, num_lines + 1):
        writer.writerow([f'E{i:05}', str(620589966 + i).zfill(10), str(620589967 + i).zfill(10), str(620589968 + i).zfill(10), '1'])

print("CSV file with 1 million lines created successfully.")