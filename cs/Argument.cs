using System;
using System.Text;
using System.Collections.Generic;
[Serializable]
public class Argument
{
    // public string title{ get; set; }
    public string description { get; set; }
	public int upvotes{ get; set; }
	public int downvotes{ get; set; }
    // public List<Motion> motionList{ get; set; }
    // public int groupSize{ get; set; }
    // public Argument()
    // {
    //     this.title = "No Title";
    //     this.description = "No Description";
    //     this.groupSize = 1;
	// 	this.motionList = new List<Motion>();

    // }
    public Argument(string description)
    {
        // this.title = title;
        this.description = description;
		this.upvotes = 0;
		this.downvotes = 0;
        // this.groupSize = 1;
		// this.motionList = new List<Motion>();

    }
	public Argument(string description, int upvotes, int downvotes)
	{
		this.description = description;
		this.upvotes = upvotes;
		this.downvotes = downvotes;
	}
    // public Argument(string title, string description)
    // {
    //     this.title = title;
    //     this.description = description;
    //     this.groupSize = 1;
	// 	this.motionList = new List<Motion>();
    // }

    // public string getTitle()
    // {
    //     return this.title;
    // }

    // public string getDescription()
    // {
    //     return this.description;
    // }

    // public int getGroupSize()
    // {
    //     return this.groupSize;
    // }
    // public void changeMotion(int motionType)
    // {
	// 	switch(motionType)
	// 	{
	// 		case 1: 
	// 			Vote vote = new Vote();
	// 			this.motionList.Add(vote);
	// 			break;
	// 		case 2:
	// 			Amend amend = new Amend("");
	// 			this.motionList.Add(amend);
	// 			break;
	// 		default:
	// 			break;	
	// 	}
    // }
    // public Motion getMotion()
    // {
    //     return this.motionList[this.motionList.Count - 1];
    // }

	// public int motionCount()
	// {
	// 	return this.motionList.Count;
	// }

	public string export()
	{
		string result = String.Format("{0}@@@{1}@@@{2}",
				this.description,
				this.upvotes,
				this.downvotes);
		return result;
	}

	//Encoder
	public byte[] encode()
	{
		byte[] data = Encoding.UTF8.GetBytes(export());
		return data;
	}
}

