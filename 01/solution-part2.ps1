function getWindowSum([int[]] $data, [int] $startIndex, [int] $windowSize = 3)
{
    $endIndex = $startIndex + $windowSize - 1
    ($data[$startIndex..$endIndex] | Measure-Object -Sum).Sum
}

[int[]] $data = Get-Content -Path 'input.txt'
$result = 0

for ($i = 0; $i -le ($data.Count - 3); $i++)
{
    if ((getWindowSum $data $i) -lt (getWindowSum $data ($i + 1)))
    {
        $result++
    }
}

$result