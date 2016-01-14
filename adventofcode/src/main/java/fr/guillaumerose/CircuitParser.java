package fr.guillaumerose;

import java.util.HashMap;
import java.util.Map;
import java.util.function.BiFunction;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import fr.guillaumerose.operators.Formula;
import fr.guillaumerose.operators.Not;
import fr.guillaumerose.operators.Number;
import fr.guillaumerose.operators.TwoVariablesFormula;

public class CircuitParser {
	private static final Pattern TWO_VARIABLES = Pattern.compile("(.*?) (.*?) (.*?)");
	private static final Pattern NOT = Pattern.compile("NOT (.*?)");
	private static final Map<String, BiFunction<Integer, Integer, Integer>> OPERATORS = operators();
	private final Map<String, Formula> formulas = new HashMap<>();

	public void parse(String input) {
		String[] split = input.split(" -> ");
		formulas.put(split[1], formula(split[0]));
	}

	public Circuit build() {
		return new Circuit(formulas);
	}

	private Formula formula(String expression) {
		Matcher matcher = NOT.matcher(expression);
		if (matcher.matches()) {
			return Not.parse(matcher.group(1));
		}
		matcher = TWO_VARIABLES.matcher(expression);
		if (matcher.matches()) {
			return TwoVariablesFormula.parse(OPERATORS.get(matcher.group(2)), matcher.group(1), matcher.group(3));
		}
		return Number.parse(expression);
	}

	private static Map<String, BiFunction<Integer, Integer, Integer>> operators() {
		Map<String, BiFunction<Integer, Integer, Integer>> op = new HashMap<>();
		op.put("AND", (a, b) -> a & b);
		op.put("OR", (a, b) -> a | b);
		op.put("LSHIFT", (a, b) -> a << b);
		op.put("RSHIFT", (a, b) -> a >> b);
		return op;
	}
}
