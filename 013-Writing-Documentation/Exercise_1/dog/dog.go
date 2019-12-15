/* This package takes human years and converts into dog years */
package dog

/* Years is an exported function, which takes human_years as a parameter and converts into dog years
   1 human_year = 7 dog_years
   input is an integer and output expected is also an integer
*/
func Years(human_years int) int {
	return human_years * 7
}
