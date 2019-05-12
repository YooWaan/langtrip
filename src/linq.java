import java.util.List;
import java.util.function.Function;
import java.util.stream.IntStream;
import java.util.stream.Collectors;

Function<IntStream,String> join = (s) -> {
	return String.join(",",
					   s.mapToObj(Integer::toString)
					   .collect(Collectors.toList()));
};

int[] data = IntStream.range(1, 10).toArray();

IntStream linq = IntStream.of(data);
System.out.println("data    : " + join.apply(linq));

linq = IntStream.of(data).filter(i -> i > 4);
System.out.println("filtered: " + join.apply(linq));

/exit
