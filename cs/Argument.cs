// using System;
// using System.IO;
public class Argument
{
    private string title;
    private string description;
    private string state;
    public Argument()
    {
        this.title = "No Title";
        this.description = "No Description";
        this.state = "Default State";
    }
    public Argument(string title)
    {
        this.title = title;
        this.description = title;
        this.state = "Initial";
    }
    public Argument(string title, string description)
    {
        this.title = title;
        this.description = description;
        this.state = "Initial";
    }

    public string getTitle()
    {
        return this.title;
    }

    public string getDescription()
    {
        return this.description;
    }
    public string getState()
    {
        return this.state;
    }
}
