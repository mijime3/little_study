<?php

function is_prime($n)
{
	for( $i = 2; $i < $n; $i++ )
	{
		if( $n % $i == 0 )
		{
			return false;
		}
	}
	return true;
}

for( $i = 2; $i <= 50000; $i++ )
{
	if( is_prime($i) ) print "[$i]\n";
}	
