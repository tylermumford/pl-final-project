using System;
using System.Text;
[Serializable]
public class Vote:Motion
{
    int numberOfVotes;
    int inFavor;
    int against;
    bool passed;
    public Vote():base(1)
    {
		this.numberOfVotes = 0;
        this.inFavor = 0;
        this.against = 0;
        this.passed = false;
    }

    public void castAgainst(int groupSize)
    {
		if(this.seconded)
		{
			this.against+=1;
        	completeArgument(groupSize);
		}
    }

    public void castInFavor(int groupSize)
    {
		if(this.seconded)
		{
			this.inFavor+=1;
        	completeArgument(groupSize);
		}
    }

    public void completeArgument(int groupSize)
    {
        if(numberOfVotes > (groupSize / 2) + 1)
        {
            if(inFavor > against)
            {
                this.passed = true;
            }
            else 
            {
                this.passed = false;
            }
        }
    }
	public string export()
	{
		string result = String.Format("NumberOfVotes:{0}\nInFavor:{1}\nAgainst:{2}\nPassed:{3}\n",
				this.numberOfVotes,
				this.inFavor,
				this.against,
				this.passed);
		return result;
	}

	//Encoder
	public byte[] encode()
	{
		byte[] data = Encoding.UTF8.GetBytes(export());
		return data;
	}
}