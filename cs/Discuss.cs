using System;
using System.Text;
using System.Collections.Generic;
[Serializable]
public class Discuss:Motion
{
    List<string> comments;

    public Discuss():base(4){}

    public void addComment(string message)
    {
		comments.Add(message);
    }

	public string export()
	{
		string result = "Comments:";
		comments.ForEach(i => result += (i+"\n\n"));
		return result;
	}

	//Encoder
	public byte[] encode()
	{
		byte[] data = Encoding.UTF8.GetBytes(export());
		return data;
	}
}