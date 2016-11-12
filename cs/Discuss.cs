using System;
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
}