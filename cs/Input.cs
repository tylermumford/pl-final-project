using System;
using System.Text;
using System.IO;
public class Input
{
    public static void Main(string[] args)
    {
        Stream inputStream = Console.OpenStandardInput();
        byte[] bytes = new byte[100];
        int outLength = inputStream.Read(bytes, 0, 100);
        char[] chars = Encoding.UTF7.GetChars(bytes, 0, outLength);
        string inStr = new string(chars);
        interpret(inStr);
    }

    // Commands: 
    // Make a new Argument with title and Description: !arg@@@Title@@@Description
    public static void interpret(string interpretString)
    {
        Console.WriteLine("interpreting...");
        string[] delims = { "@@@" };
        string[] args = interpretString.Split(delims, StringSplitOptions.RemoveEmptyEntries);
        switch (args[0])
        {
            case "!arg":
                Argument argument = new Argument(args[1], args[2]);
                Console.WriteLine("Title: " + argument.getTitle() + "," + "Description: " + argument.getDescription());
                break;
            default:
                break;
        }
    }
}
