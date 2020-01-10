import sys
import re

#Jonathan Burdette
#Text Cloud Implementation

if len(sys.argv) < 4:
    print("Invalid syntax. To run program, enter: \"python text_cloud.py <input file name> <exclude file name> <output file name>\"")
else:
    input_file = sys.argv[1]
    exclude_file = sys.argv[2]
    output_file = sys.argv[3]

    #store words from exclude file in dict (blank values)
    exclude_dict = {}
    f = open(exclude_file, "r")
    lines = f.readlines()
    for line in lines:
        words = re.split("[^A-Z'a-z]+", line)
        for word in words:
            exclude_dict[word] = ""

    #store the words and their counts from input file
    input_dict = {}
    f = open(input_file, "r", encoding="utf8")
    lines = f.readlines()
    for line in lines:
        words = re.split("[^A-Z'a-z]+", line)
        for word in words:
            word = word.lower()
            if exclude_dict.get(word) == None and re.match("[a-zA-Z]", word) and len(word) > 1:
                count = input_dict.get(word)
                if count == None:
                    input_dict[word] = 1
                else:
                    input_dict[word] = count + 1
               
    #sort the words based on count
    sorted_counts = sorted(input_dict.items(), reverse=True, key=lambda x: x[1])

    #take the top fifty words based on count
    top_fifty = sorted_counts[0:50]

    #sort the top fifty in alphabetical order
    top_fifty.sort()

    #info for HTML file
    range = (sorted_counts[0][1] - sorted_counts[49][1])
    size_factor = 1000 / range
    colors = ["#623224", "#FABE8B", "#EC1D23", "#0160B0"]

    f = open(output_file, "w") 
    color_count = 0
    for var in top_fifty:
        if color_count == 4: #colors restart after 4 iterations
            color_count = 0

        fontsize = var[1] * size_factor
        f.write("<span style=\"font-size:")
        f.write(str(fontsize))
        f.write("%; color:")
        f.write(colors[color_count])
        f.write(";\">")
        f.write(var[0])
        f.write("</span> &nbsp; &nbsp;\n")
        color_count = color_count + 1