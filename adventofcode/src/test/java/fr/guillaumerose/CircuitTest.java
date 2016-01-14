package fr.guillaumerose;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.Test;

public class CircuitTest {
	private final CircuitParser parser = new CircuitParser();

	@Test
	public void should_load_hard_coded_value_signal() throws Exception {
		parser.parse("123 -> x");
		assertThat(parser.build().valueFor("x")).isEqualTo(123);
	}

	@Test
	public void should_copy_value_signal() throws Exception {
		parser.parse("x -> y");
		parser.parse("123 -> x");
		assertThat(parser.build().valueFor("y")).isEqualTo(123);
	}

	@Test
	public void should_load_hard_several_signals() throws Exception {
		parser.parse("123 -> x");
		parser.parse("456 -> y");
		assertThat(parser.build().valueFor("y")).isEqualTo(456);
		assertThat(parser.build().valueFor("x")).isEqualTo(123);
	}

	@Test
	public void should_support_complement() throws Exception {
		parser.parse("7 -> x");
		parser.parse("NOT x -> y");
		assertThat(parser.build().valueFor("y")).isEqualTo(65528);
	}

	@Test
	public void should_support_complement_with_hard_coded_value() throws Exception {
		parser.parse("NOT 2 -> x");
		assertThat(parser.build().valueFor("x")).isEqualTo(65533);
	}

	@Test
	public void should_apply_and_to_signals() throws Exception {
		parser.parse("3 -> x");
		parser.parse("x AND 2 -> y");
		assertThat(parser.build().valueFor("y")).isEqualTo(2);
	}

	@Test
	public void should_apply_or_to_signals() throws Exception {
		parser.parse("1 -> x");
		parser.parse("2 OR x -> y");
		assertThat(parser.build().valueFor("y")).isEqualTo(3);
	}

	@Test
	public void should_apply_lshift_to_signal() throws Exception {
		parser.parse("1 -> x");
		parser.parse("x LSHIFT 2 -> y");
		assertThat(parser.build().valueFor("y")).isEqualTo(4);
	}

	@Test
	public void should_apply_rshift_to_signal() throws Exception {
		parser.parse("8 RSHIFT 2 -> y");
		assertThat(parser.build().valueFor("y")).isEqualTo(2);
	}

	@Test
	public void should_validate_example() throws Exception {
		parser.parse("123 -> x");
		parser.parse("456 -> y");
		parser.parse("x AND y -> d");
		parser.parse("x OR y -> e");
		parser.parse("x LSHIFT 2 -> f");
		parser.parse("y RSHIFT 2 -> g");
		parser.parse("NOT x -> h");
		parser.parse("NOT y -> i");

		assertThat(parser.build().valueFor("d")).isEqualTo(72);
		assertThat(parser.build().valueFor("e")).isEqualTo(507);
		assertThat(parser.build().valueFor("f")).isEqualTo(492);
		assertThat(parser.build().valueFor("g")).isEqualTo(114);
		assertThat(parser.build().valueFor("h")).isEqualTo(65412);
		assertThat(parser.build().valueFor("i")).isEqualTo(65079);
	}

	@Test
	public void should_handle_not_already_defined_variable() throws Exception {
		parser.parse("2 AND x -> z");
		parser.parse("3 -> x");
		assertThat(parser.build().valueFor("z")).isEqualTo(2);
	}
}
