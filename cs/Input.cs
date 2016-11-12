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
		var dataFolder = "/vagrant/data/storage/";
        try
        {
            string[] delims = { "@@@" };
            string[] args = interpretString.Split(delims, StringSplitOptions.RemoveEmptyEntries);
            string fileName = dataFolder + args[0] + ".txt";
			Argument argue = null;
            if(!args[1].Equals("!arg"))
			{
				argue = deArgument(fileName);
			}
			using (FileStream fs = File.Open(fileName, FileMode.OpenOrCreate, FileAccess.Write))
            {
				// IFormatter formatter = new BinaryFormatter();
                switch (args[1])
                {
                    case "!arg":
                        Argument argueDef = new Argument(args[2]);
						fs.Write(argueDef.encode(),0,argueDef.encode().Length);
                        break;
                    // case "!vote":
					// 	argue.changeMotion(1);
					// 	fs.Write((argue.getMotion() as Vote).encode(),0,(argue.getMotion() as Vote).encode().Length);
					// 	break;
					case "!upvote":
						argue.upvotes += 1;
						fs.Write(argue.encode(),0,argue.encode().Length);
						break;
					case "!downvote":
						argue.downvotes += 1;
						fs.Write(argue.encode(),0,argue.encode().Length);
						break;
                    // case "!discuss":
					// 	argue.changeMotion(4);
					// 	fs.Write((argue.getMotion() as Discuss).encode(),0,(argue.getMotion() as Discuss).encode().Length);
					// 	break;
					// case "!secondcurrentmotion":
					// 	// System.Console.WriteLine(argue3.getTitle());
					// 	// System.Console.WriteLine(argue3.motionCount());
					// 	// System.Console.WriteLine(argue3.getMotion().getMotionTitle());
					// 	argue.getMotion().secondMotion();
					// 	// formatter.Serialize(fs, argue3);
					// 	break;
					// case "!amend":
					// 	break;
                    // case "!table":
					// 	argue.changeMotion(3);
					// 	// formatter.Serialize(fs, argue4);
					// 	break;
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

	public static Argument deArgument(string path)
	{
		// IFormatter formatter = new BinaryFormatter();
		string text = File.ReadAllText(path);
		string[] delims = { "@@@" };
        string[] args = text.Split(delims, StringSplitOptions.RemoveEmptyEntries);
		Argument argue = new Argument(args[0], Convert.ToInt32(args[1]), Convert.ToInt32(args[2]));
		return argue;
	}
}
