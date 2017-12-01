package fr.guillaumerose.operators;

import static fr.guillaumerose.Lists.newArrayList;
import static fr.guillaumerose.Numbers.isNumeric;

import java.util.List;

public class Not implements Formula {
	private final String variable;

	public Not(String input) {
		variable = input;
	}

	@Override
	public List<String> variables() {
		return newArrayList(variable);
	}

	@Override
	public Integer compute(List<Integer> inputs) {
		return not(inputs.get(0));
	}

	private static int not(Integer integer) {
		return ~integer & 0xFFFF;
	}

	public static Formula parse(String input) {
		if (isNumeric(input)) {
			return new Number(not(Integer.valueOf(input)));
		}
		return new Not(input);
	}
}