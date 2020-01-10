import java.io.*;
import java.util.*;

//Jonathan Burdette
//Text Cloud Implementation

public class TextCloud {
	
	private HashMap<String, String> excludeHM = new HashMap<String, String>(150);
	private HashMap<String, Integer> inputHM = new HashMap<String, Integer>(30000);
	private List<Map.Entry<String, Integer>> sortedCounts;
	private TreeMap<String, Integer> finalTM = new TreeMap<String, Integer>();

	//read in the exclude file
	private void readExclude(String excludeFile) {
		try {
			File file = new File(excludeFile);
			Scanner sc = new Scanner(file);
			
			//store words from exclude file in hash map (blank values)
			while(sc.hasNextLine()) {
				excludeHM.put(sc.nextLine(), "");
			}
			sc.close();
			
		} catch(IOException e) {
			System.out.println("Error reading file");
		}
	}
	
	//read the input file and store for further use
	private void readInput(String inputFile) {
		try {
			File file = new File(inputFile);
			Scanner sc = new Scanner(file);
			sc.useDelimiter("[^a-z'A-Z]+");
			
			//store words and counts where key is string and value is an integer used to hold the count
			while(sc.hasNext()) {
				String word = sc.next().toLowerCase();
				if(!excludeHM.containsKey(word)) { //don't allow excluded words
				   if(!word.matches("[a-z]") && !word.matches("'")) { //since a and i are in exclude file it is fine to filter them out
						Integer count = inputHM.get(word);
						if(count == null) {
							inputHM.put(word, 1);
						} else {
							inputHM.put(word, count + 1);
						}
				   }
				}
			}
			sc.close();
			
		} catch(IOException e) {
			System.out.println("Error reading file");
		}
	}
	
	//find the 50 most common words based on counts
	private void findCommonWords() {
		
		//create list from hash map
		List<Map.Entry<String, Integer>> inputList = new LinkedList<Map.Entry<String, Integer>>(inputHM.entrySet());
		
		//create a comparator to compare values from input list
		Comparator<Map.Entry<String, Integer>> comp = new Comparator<Map.Entry<String, Integer>>() { 
			public int compare(Map.Entry<String, Integer> count1, Map.Entry<String, Integer> count2) {
				return count1.getValue().compareTo(count2.getValue());
			}
		};
		
		//sort the list using comparator
		Collections.sort(inputList, comp);
				
		//reverse list for easy access to top 50
		Collections.reverse(inputList);
		
		//create new list using top 50
		sortedCounts = new LinkedList<Map.Entry<String, Integer>>(inputList.subList(0, 50));
		
		//sort new list by key by inserting into tree map
		for(Map.Entry<String, Integer> entry : sortedCounts) {
			finalTM.put(entry.getKey(), entry.getValue());
		}		
	}
	
	//writes html for text cloud
	private void writeHTMLFile(String htmlFile) {
		
		//info for html file
		Map.Entry<String, Integer> highest = sortedCounts.get(0);
		Map.Entry<String, Integer> lowest = sortedCounts.get(49);
		int range = (highest.getValue() - lowest.getValue());
		int sizeFactor = 1000 / range;				
		String[] colors = {"#008000", "#800080", "#7D542A", "#2B547E"};
		
		try {
			BufferedWriter bw = new BufferedWriter(new FileWriter(htmlFile));
			int count = 0;
			for(Map.Entry<String, Integer> entry : finalTM.entrySet()) {
				if(count == 4) { //colors restart after 4 iterations
					count = 0;
				}
				int fontSize = entry.getValue() * sizeFactor;
				bw.write("<span style=\"font-size:"+fontSize+"%; color:"+colors[count]+";\">"+entry.getKey()+"</span> &nbsp; &nbsp;");
				count++;
			}
			bw.close();
			
		} catch (IOException e) {
			System.out.println("Error reading file");
		}
	}
	
	public static void main(String[] args) {
		
		TextCloud tc = new TextCloud();
		
		//check command line syntax
		if(args.length < 3) {
			System.out.println("Invalid syntax. To run program, enter: \"java TextCloud <input file name> <exclude file name> <output file name>\"");
		} else {
			tc.readExclude(args[1]);
			tc.readInput(args[0]);
			tc.findCommonWords();
			tc.writeHTMLFile(args[2]);
		}
	}
}
