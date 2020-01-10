#Jonathan Burdette
#Text Cloud Implementation

if ARGV.length < 3
    puts "Invalid syntax. To run program, enter: \"ruby text_cloud.rb <input file name> <exclude file name> <output file name>\""
    exit
else
    input_file = ARGV[0]
    exclude_file = ARGV[1]
    output_file = ARGV[2]
    
    #store words from exclude file in hash (blank values)
    exclude_hash = Hash.new
    f = File.open(exclude_file, "r")
    f.each do |line|
        words = line.split(/[^A-Z'a-z]/)
        words.each do |word|
            exclude_hash[word] = ""
        end
    end

    #store the words and their counts from input file
    input_hash = Hash.new
    f = File.open(input_file, "r")
    f.each do |line|
        words = line.split(/[^A-Z'a-z]/)
        words.each do |word|
            word = word.downcase
            if !exclude_hash.has_key?(word) && word.match(/[A-Za-z]/) && word.length > 1 
                count = input_hash[word]
                if count == nil
                    input_hash[word] = 1
                else
                    input_hash[word] = count + 1
                end
            end
        end
    end

    #sort the words based on count
    sorted_counts = input_hash.sort_by {|key, value| -value}

    #take the top fifty words based on count
    top_fifty = sorted_counts[0, 50]

    #sort the top fifty in alphabetical order
    sorted_words = top_fifty.sort_by {|key, value| -key}

    #info for HTML file
    range = (top_fifty[0][1] - top_fifty[49][1]).to_f
    size_factor = (1000 / range).to_f
    colors = ["#623224", "#FABE8B", "#EC1D23", "#0160B0"]

    f = File.open(output_file, "w") 
    color_count = 0
    sorted_words.each do |key, value|
        if color_count == 4 #colors restart after 4 iterations
            color_count = 0
        end
        fontsize = "#{value}".to_f * size_factor
        f.print "<span style=\"font-size:"
        f.print fontsize
        f.print "%; color:"
        f.print colors[color_count]
        f.print ";\">"
        f.print "#{key}"
        f.print "</span> &nbsp; &nbsp;\n"
        color_count = color_count + 1
    end
end