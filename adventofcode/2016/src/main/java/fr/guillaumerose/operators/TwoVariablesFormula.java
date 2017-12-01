package fr.guillaumerose.operators;

import static fr.guillaumerose.Lists.newArrayList;
import static fr.guillaumerose.Numbers.isNumeric;
import static java.lang.Integer.valueOf;

import java.util.List;
import java.util.function.BiFunction;
import java.util.function.Function;

public class TwoVariablesFormula implements Formula {
	private final List<String> variables;
	private final Function<List<Integer>, Integer> fun;

	private TwoVariablesFormula(Function<List<Integer>, Integer> fun, String... variables) {
		this.variables = newArrayList(variables);
		this.fun = fun;
	}

	@Override
	public List<String> variables() {
		return variables;
	}

	@Override
	public Integer compute(List<Integer> values) {
		return fun.apply(values);
	}

	public static Formula parse(BiFunction<Integer, Integer, Integer> fun, String input1, String input2) {
		if (isNumeric(input1) && isNumeric(input2)) {
			return new TwoVariablesFormula(list -> fun.apply(valueOf(input1), valueOf(input2)));
		} else if (isNumeric(input1) && !isNumeric(input2)) {
			return new TwoVariablesFormula(list -> fun.apply(valueOf(input1), list.get(0)), input2);
		} else if (!isNumeric(input1) && isNumeric(input2)) {
			return new TwoVariablesFormula(list -> fun.apply(list.get(0), valueOf(input2)), input1);
		}
		return new TwoVariablesFormula(list -> fun.apply(list.get(0), list.get(1)), input1, input2);
	}
}