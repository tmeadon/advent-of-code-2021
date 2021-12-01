[int[]] $data = Get-Content -Path 'input.txt'
$result = 0

for ($i = 1; $i -le $data.Count; $i++)
{
    if ($data[$i] -gt $data[$i - 1])
    {
        $result++
    }
}

$result