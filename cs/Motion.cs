using System;
[Serializable]
public enum MotionType
{
    Initial = 0,
    Vote = 1,
    Amend = 2,
    Table = 3,
    Discuss = 4
}
[Serializable]
public class Motion
{
    private MotionType type{ get; set; }
    private string title{ get; set; }
    public bool seconded{  get;  set; }
    private bool canceled{ get; set; }
    public Motion(int state)
    {
        type = (MotionType)state;
        switch (type)
        {
            case MotionType.Initial:
                this.title = "Initial";
                break;
            case MotionType.Vote:
                this.title = "Vote";
                break;
            case MotionType.Amend:
                this.title = "Amend";
                break;
            case MotionType.Table:
                this.title = "Table";
                break;
            case MotionType.Discuss:
                this.title = "Discuss";
                break;
            default:
                this.title = "Default";
                break;
        }
        this.seconded = false;
        this.canceled = false;
    }

    public string getMotionTitle()
    {
        return this.title;
    }
    public void secondMotion()
    {
        if (seconded == false)
        {
            this.seconded = true;
        }
    }

    public void cancel()
    {
        this.canceled = true;
    }

	public int getMotionType()
	{
		return (int)type;
	}
}