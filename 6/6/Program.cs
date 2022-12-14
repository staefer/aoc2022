string text = System.IO.File.ReadAllText("input.txt");

//Console.WriteLine(findStartPacketMarker(text));

Console.WriteLine(findStartPacketMarker(text));

int findStartPacketMarker(string input) {
    int pos = 0;
    string values;

    while (true) {
        // For part 1
        // values = input.Substring(pos, 4)
        // For part 2
        values = input.Substring(pos, 14);
        Console.WriteLine("Substring: {0}", values);
        if (!hasDuplicates(values)) {
            Console.WriteLine("Duplicate found! Substring: {0} Position {1}", values, pos);
            // For part 1
            // return pos + 4
            // For part 2
            return pos + 14;
        } 
        else {
            pos++;
        }
    }
}

bool hasDuplicates(string values) {
    for (int i = 0; i < values.Length - 1; i++) {
        for (int j = i + 1; j < values.Length; j++) {
            if (values[i].Equals(values[j])) {
                return true;
            }
        }
    }
    return false;
}