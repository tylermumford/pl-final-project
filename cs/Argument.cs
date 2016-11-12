using System;
using System.Collections.Generic;
[Serializable]
public class Argument
{
    private string title{ get; set; }
    private string description { get; set; }
    private List<Motion> motionList;
    private int groupSize{ get; set; }
    public Argument()
    {
        this.title = "No Title";
        this.description = "No Description";
        this.groupSize = 1;
		this.motionList = new List<Motion>();

    }
    public Argument(string title)
    {
        this.title = title;
        this.description = title;
        this.groupSize = 1;
		this.motionList = new List<Motion>();

    }
    public Argument(string title, string description)
    {
        this.title = title;
        this.description = description;
        this.groupSize = 1;
		this.motionList = new List<Motion>();
    }

    public string getTitle()
    {
        return this.title;
    }

    public string getDescription()
    {
        return this.description;
    }

    public int getGroupSize()
    {
        return this.groupSize;
    }
    public void changeMotion(int motionType)
    {
		switch(motionType)
		{
			case 1: 
				Vote vote = new Vote();
				this.motionList.Add(vote);
				break;
			case 2:
				Amend amend = new Amend("");
				this.motionList.Add(amend);
				break;
			default:
				break;	
		}
    }
    public Motion getMotion()
    {
        return this.motionList[this.motionList.Count - 1];
    }

	public int motionCount()
	{
		return motionList.Count;
	}
}

