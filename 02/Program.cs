var inputFile = "input.txt";
var part1Sub = new Part1Sub();
var part2Sub = new Part2Sub();

foreach (var instruction in File.ReadAllLines(inputFile))
{
    part1Sub.ApplyInstruction(instruction);
    part2Sub.ApplyInstruction(instruction);
}

Console.WriteLine($"Part 1 result = {part1Sub.ForwardPosition * part1Sub.Depth}");
Console.WriteLine($"Part 2 result = {part2Sub.ForwardPosition * part2Sub.Depth}");