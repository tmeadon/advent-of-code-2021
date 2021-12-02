abstract class Sub
{
    public int ForwardPosition { get; protected set; }
    public int Depth { get; protected set; }

    public void ApplyInstruction(string instruction)
    {
        var direction = instruction.Split(' ')[0];
        var amount = int.Parse(instruction.Split(' ')[1]);

        switch (direction)
        {
            case "forward":
                Forward(amount);
                break;

            case "up": 
                Up(amount);
                break;

            case "down":
                Down(amount);
                break;
        }
    }

    protected abstract void Down(int amount);
    protected abstract void Up(int amount);
    protected abstract void Forward(int amount);
}

class Part1Sub : Sub
{
    protected override void Down(int amount) => Depth += amount;

    protected override void Up(int amount) => Depth -= amount;

    protected override void Forward(int amount) => ForwardPosition += amount;
}

class Part2Sub : Sub
{
    public int Aim { get; protected set; } = 0;

    protected override void Down(int amount) => Aim += amount;

    protected override void Up(int amount) => Aim -= amount;

    protected override void Forward(int amount)
    {
        ForwardPosition += amount;
        Depth += Aim * amount;
    }
}