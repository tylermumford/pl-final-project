public class Amend:Motion
{
    string amendment;
    public Amend(string amendment):base(2)
    {
        this.amendment = amendment;
    }

    public string getAmendment()
    {
        return amendment;
    }
}