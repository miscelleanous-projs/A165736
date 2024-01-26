import std.stdio : writeln;
import std.range : iota;
import std.algorithm : map,sum;
import std.bigint : BigInt, powmod;

// OEIS-A165736
// https://oeis.org/A165736

BigInt a165736(BigInt n)
{
    BigInt modVal = BigInt("10_000_000_000");

    BigInt result = n;

    for (int i = 0; i < 10; i++)
        result = powmod(n, result, modVal);

    return result;
}

void main()
{
    iota(1.BigInt, 100.BigInt + 1).map!a165736.writeln;
}

unittest
{
    assert(a165736(1.BigInt) == 1.BigInt);
    //
    assert(a165736( 10.BigInt) == 0.BigInt);
    assert(a165736( 20.BigInt) == 0.BigInt);
    assert(a165736( 30.BigInt) == 0.BigInt);
    assert(a165736( 40.BigInt) == 0.BigInt);
    assert(a165736(100.BigInt) == 0.BigInt);
    //
    assert(a165736( 6.BigInt) == 7_447_238_656.BigInt);
    assert(a165736(11.BigInt) == 9_172_666_611.BigInt);
    assert(a165736(16.BigInt) ==   290_415_616.BigInt);
    assert(a165736(19.BigInt) ==   609_963_179.BigInt);
    assert(a165736(23.BigInt) == 1_075_718_247.BigInt);
}