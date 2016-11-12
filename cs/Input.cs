using System;
using System.Text;
using System.IO;
using System.Collections.Generic;

public class Input
{
    public static void Main(string[] args)
    {
		interpret(args);
    }

    // Commands: 
    // First put "Filename###" for which Argument to edit
    // Make a new Argument with title and Description: !arg@@@Title@@@Description
    public static void interpret(string[] args)
    {
		// FOR DEBUGGING ARGUMENTS
		// foreach (var item in args)
		// {
		// 	System.Console.WriteLine("Argument: {0}", item);
		// }

		var dataFolder = "/vagrant/data/storage/";
        try
        {
            string fileName = dataFolder + args[0] + ".txt";
			Argument argue = null;
            if(!args[1].Equals("create"))
			{
				argue = deArgument(fileName);
			}
			using (FileStream fs = File.Open(fileName, FileMode.OpenOrCreate, FileAccess.Write))
            {
                switch (args[1])
                {
                    case "create":
                        Argument argueDef = new Argument(args[2]);
						fs.Write(argueDef.encode(),0,argueDef.encode().Length);
						Console.Error.WriteLine(argueDef.export());
                        break;
                    // case "!vote":
					// 	argue.changeMotion(1);
					// 	fs.Write((argue.getMotion() as Vote).encode(),0,(argue.getMotion() as Vote).encode().Length);
					// 	break;
					case "upvote":
						argue.upvotes += 1;
						fs.Write(argue.encode(),0,argue.encode().Length);
						Console.WriteLine(argue.export());
						break;
					case "downvote":
						argue.downvotes += 1;
						fs.Write(argue.encode(),0,argue.encode().Length);
						Console.WriteLine(argue.export());
						break;
					case "export":
						System.Console.WriteLine(argue.export());	
						break;
                    // case "!discuss":
					// 	argue.changeMotion(4);
					// 	fs.Write((argue.getMotion() as Discuss).encode(),0,(argue.getMotion() as Discuss).encode().Length);
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
		string text = File.ReadAllText(path);
		string[] delims = { "@@@" };
        string[] args = text.Split(delims, StringSplitOptions.RemoveEmptyEntries);
		Argument argue = new Argument(args[0], Convert.ToInt32(args[1]), Convert.ToInt32(args[2]));
		return argue;
	}
}
