using System;
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

    public void castAgainst()
    {
		if(this.seconded)
		{
			this.against+=1;
        	completeArgument();
		}
    }

    public void castInFavor()
    {
		if(this.seconded)
		{
			this.inFavor+=1;
        	completeArgument();
		}
    }

    public void completeArgument()
    {
        if(numberOfVotes > (this.getGroupSize() / 2) + 1)
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
}