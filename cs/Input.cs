using System;
using System.Text;
using System.IO;
using System.Runtime.Serialization;
using System.Runtime.Serialization.Formatters.Binary;
public class Input
{
    public static void Main(string[] args)
    {
        Stream inputStream = Console.OpenStandardInput();
        byte[] bytes = new byte[100];
        int outLength = inputStream.Read(bytes, 0, 100);
        char[] chars = Encoding.UTF8.GetChars(bytes, 0, outLength);
        string inStr = new string(chars);
        interpret(inStr);




        // using(Stream stream = new FileStream("Arg.txt", FileMode.Create, FileAccess.Write, FileShare.None)){
        // 	Argument argue = new Argument("Trump", "Trump v Hilary");
        // 	IFormatter formatter = new BinaryFormatter();
        // 	formatter.Serialize(stream, argue);
        // }


        // using(Stream stream1 = new FileStream("Arg.txt", FileMode.Open, FileAccess.Read, FileShare.Read)){
        // 	IFormatter formatter1 = new BinaryFormatter();
        // 	Argument argue1 = (Argument) formatter1.Deserialize(stream1);
        // 	System.Console.WriteLine(argue1.getTitle());
        // 	System.Console.WriteLine(argue1.getDescription());
        // }

    }

    // Commands: 
    // First put "Arugment###" for which Argument to edit
    // Make a new Argument with title and Description: !arg@@@Title@@@Description
    public static void interpret(string interpretString)
    {
        try
        {
            string[] delims = { "@@@" };
            string[] args = interpretString.Split(delims, StringSplitOptions.RemoveEmptyEntries);
            string fileName = args[0] + ".txt";
            using (FileStream fs = File.Open(fileName, FileMode.OpenOrCreate, FileAccess.ReadWrite))
            {
				IFormatter formatter = new BinaryFormatter();
                switch (args[1])
                {
                    case "!arg":
                        Argument argue = new Argument(args[2], args[3]);
						formatter.Serialize(fs, argue);
                        break;
                    case "!vote":
						Argument argue1 = deArgument(fs);
						argue1.changeMotion(1);
						formatter.Serialize(fs, argue1);
						break;
                    case "!discuss":
						Argument argue2 = deArgument(fs);
						argue2.changeMotion(4);
						Discuss discuss = (Discuss) argue2.getMotion();
						formatter.Serialize(fs, argue2);
						break;
					case "!secondcurrentmotion":
						Argument argue3 = deArgument(fs);
						System.Console.WriteLine(argue3.motionCount());
						System.Console.WriteLine(argue3.getMotion().getMotionTitle());
						argue3.getMotion().secondMotion();
						formatter.Serialize(fs, argue3);
						break;
					case "!amend":
						break;
                    case "!table":
						Argument argue4 = deArgument(fs);
						argue4.changeMotion(3);
						Table table = (Table) argue4.getMotion();
						formatter.Serialize(fs, argue4);
						break;
                    default:
						System.Console.WriteLine("default switch");
                        break;
                }
            }
        }
        catch (Exception ex)
        {
            System.Console.WriteLine(ex.ToString());
        }
    }

	public static Argument deArgument(FileStream fs)
	{
		IFormatter formatter = new BinaryFormatter();
		return (Argument) formatter.Deserialize(fs);
	}
}
